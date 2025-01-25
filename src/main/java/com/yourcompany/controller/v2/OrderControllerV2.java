@RestController
@RequestMapping("/api/v2/orders")
@RequiredArgsConstructor
public class OrderControllerV2 {

    private final OrderService orderService;
    private final PaymentService paymentService;

    @PostMapping
    @CircuitBreaker(name = "orderService", fallbackMethod = "fallbackCreateOrder")
    @Operation(
        summary = "Create new order (v2)", 
        description = "Creates an order and processes payment."
    )
    @ApiResponses({
        @ApiResponse(responseCode = "200", description = "Order created successfully"),
        @ApiResponse(responseCode = "503", description = "Service unavailable due to fallback")
    })
    public ResponseEntity<Order> createOrderV2(@Valid @RequestBody OrderRequest request) {
        Order order = orderService.createOrder(request);
        paymentService.processPayment(order);
        return ResponseEntity.ok(order);
    }

    private ResponseEntity<Order> fallbackCreateOrder(OrderRequest request, Throwable t) {
        log.error("Failed to create order: {}", t.getMessage());
        return ResponseEntity.status(HttpStatus.SERVICE_UNAVAILABLE)
            .header("X-Fallback-Reason", t.getMessage())
            .build();
    }
}
