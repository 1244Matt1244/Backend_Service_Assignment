@Service
@RequiredArgsConstructor
public class PaymentService {

    @CircuitBreaker(name = "paymentService", fallbackMethod = "fallbackPayment")
    public void processPayment(OrderDTO order) {
        // Simulate payment processing
        if (order.quantity() > 100) {
            throw new PaymentException("Quantity exceeds limit");
        }
        // Process payment logic (e.g., call external service)
    }

    public void fallbackPayment(OrderDTO order, Throwable t) {
        logPaymentFailure(order, t);
    }

    private void logPaymentFailure(OrderDTO order, Throwable t) {
        // Log payment failure to a separate service or dead-letter queue
        System.out.printf("Payment failed for order %s (quantity: %d). Reason: %s%n",
                order.product(), order.quantity(), t.getMessage());
    }
}
