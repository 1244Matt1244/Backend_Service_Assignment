package com.example.advancedapp;

import io.github.resilience4j.circuitbreaker.annotation.CircuitBreaker;
import io.micrometer.core.annotation.Timed;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.oauth2.jwt.Jwt;
import org.springframework.web.bind.annotation.*;

import javax.persistence.*;
import javax.servlet.http.HttpServletRequest;
import java.util.Map;
import java.util.UUID;

@SpringBootApplication
public class LevelBackendApp {
    public static void main(String[] args) {
        SpringApplication.run(SeniorLevelBackendApp.class, args);
    }
}

@Entity
@Table(name = "orders")
class Order {
    @Id
    @GeneratedValue
    private UUID id;
    private String product;
    private Integer quantity;
    
    // Getters and setters
}

record OrderDTO(String product, Integer quantity) {}

interface OrderRepository extends JpaRepository<Order, UUID> {}

@RestController
@RequestMapping("/api/v1/orders")
class OrderController {
    private final OrderRepository repository;
    private final Tracer tracer;
    private final PaymentService paymentService;

    public OrderController(OrderRepository repository, Tracer tracer, PaymentService paymentService) {
        this.repository = repository;
        this.tracer = tracer;
        this.paymentService = paymentService;
    }

    @PostMapping
    @PreAuthorize("hasAuthority('SCOPE_orders:write')")
    @Timed(value = "create_order.time", description = "Time taken to create order")
    public ResponseEntity<Order> createOrder(
            @RequestBody OrderDTO orderDTO,
            @AuthenticationPrincipal Jwt jwt,
            HttpServletRequest request
    ) {
        Span span = tracer.spanBuilder("CreateOrder").startSpan();
        try (var scope = span.makeCurrent()) {
            // Business logic
            Order order = new Order();
            order.setProduct(orderDTO.product());
            order.setQuantity(orderDTO.quantity());
            
            // Process payment
            paymentService.processPayment(orderDTO);
            
            // Audit trail
            System.out.printf("User %s created order from IP %s%n",
                    jwt.getSubject(),
                    request.getRemoteAddr());
            
            return ResponseEntity.ok(repository.save(order));
        } finally {
            span.end();
        }
    }

    @GetMapping("/{id}")
    @CircuitBreaker(name = "orderService", fallbackMethod = "fallbackGetOrder")
    public ResponseEntity<Order> getOrder(@PathVariable UUID id) {
        return ResponseEntity.ok(repository.findById(id)
                .orElseThrow(() -> new OrderNotFoundException(id)));
    }

    private ResponseEntity<Order> fallbackGetOrder(UUID id, Throwable t) {
        return ResponseEntity.status(503)
                .header("X-Fallback-Reason", t.getMessage())
                .build();
    }
}

@Service
class PaymentService {
    @CircuitBreaker(name = "paymentService", fallbackMethod = "fallbackPayment")
    public void processPayment(OrderDTO order) {
        // External payment processing simulation
        if (order.quantity() > 100) {
            throw new PaymentException("Quantity exceeds limit");
        }
    }

    public void fallbackPayment(OrderDTO order, Throwable t) {
        // Log to dead letter queue
        System.out.printf("Payment fallback for order %s: %s%n", 
                order.product(), t.getMessage());
    }
}

@ControllerAdvice
class GlobalExceptionHandler {
    @ExceptionHandler(OrderNotFoundException.class)
    public ResponseEntity<Map<String, String>> handleOrderNotFound(OrderNotFoundException ex) {
        return ResponseEntity.status(404)
                .body(Map.of("error", ex.getMessage()));
    }

    @ExceptionHandler(PaymentException.class)
    public ResponseEntity<Map<String, String>> handlePaymentError(PaymentException ex) {
        return ResponseEntity.status(402)
                .body(Map.of("error", "Payment processing failed", 
                             "detail", ex.getMessage()));
    }
}

class OrderNotFoundException extends RuntimeException {
    OrderNotFoundException(UUID id) {
        super("Order not found: " + id);
    }
}

class PaymentException extends RuntimeException {
    PaymentException(String message) {
        super(message);
    }
}
