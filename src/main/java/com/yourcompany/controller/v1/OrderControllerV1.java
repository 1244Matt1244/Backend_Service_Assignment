@RestController
@RequestMapping("/api/v1/orders")
@RequiredArgsConstructor
public class OrderControllerV1 {

    private final OrderService orderService;

    @GetMapping("/{id}")
    @Operation(
        summary = "Get order (v1)", 
        deprecated = true, 
        description = "Use /api/v2/orders/{id} instead"
    )
    public ResponseEntity<Order> getOrderV1(@PathVariable String id) {
        return ResponseEntity.ok(orderService.getOrder(id));
    }
}
