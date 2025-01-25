class OrderNotFoundException extends RuntimeException {
    public OrderNotFoundException(UUID id) {
        super("Order not found: " + id);
    }
}

class PaymentException extends RuntimeException {
    public PaymentException(String message) {
        super(message);
    }
}
