@RestController
@RequestMapping("/api/v2/orders")
@RequiredArgsConstructor
public class OrderControllerV2 {

    private final OrderService orderService;
    private final PaymentService paymentService;

    @PostMapping
    @CircuitBreaker(name = "orderService", fallbackMethod = "fallbackCreateOrder")
    public ResponseEntity<Order> createOrderV2(@Valid @RequestBody OrderRequest request) {
        Order order = orderService.createOrder(request);
        paymentService.processPayment(order);
        return ResponseEntity.ok(order);
    }

    private ResponseEntity<Order> fallbackCreateOrder(OrderRequest request, Throwable t) {
        return ResponseEntity.status(HttpStatus.SERVICE_UNAVAILABLE)
            .header("X-Fallback-Reason", t.getMessage())
            .build();
    }
}
