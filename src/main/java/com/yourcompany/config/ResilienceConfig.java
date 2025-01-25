@Configuration
public class ResilienceConfig {

    @Bean
    public CircuitBreakerConfigCustomizer orderServiceCircuitBreaker() {
        return CircuitBreakerConfigCustomizer
            .of("orderService", builder -> builder
                .failureRateThreshold(50)
                .waitDurationInOpenState(Duration.ofSeconds(30))
                .slidingWindowSize(10)
                .minimumNumberOfCalls(5));
    }

    @Bean
    public RetryConfigCustomizer orderServiceRetry() {
        return RetryConfigCustomizer
            .of("orderService", builder -> builder
                .maxAttempts(3)
                .intervalFunction(IntervalFunction.ofExponentialBackoff(500, 2)));
    }
}
