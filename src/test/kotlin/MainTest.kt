import org.junit.jupiter.api.Test

public class MainTest {
    @Test
    fun parseOptions() {
        Hello().parse(arrayOf("--count=3"))
    }
}
