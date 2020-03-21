import com.github.ajalt.clikt.core.CliktCommand
import com.github.ajalt.clikt.parameters.options.default
import com.github.ajalt.clikt.parameters.options.option
import com.github.ajalt.clikt.parameters.types.int

class Hello : CliktCommand() {
    private val count: Int by option(help = "Number of greetings").int().default(1)
    private val name: String by option(help = "The person to greet").default("Casey")

    override fun run() {
        repeat(count) {
            echo("Hello $name!")
        }
    }
}

fun main(args: Array<String>) = Hello().main(args)
