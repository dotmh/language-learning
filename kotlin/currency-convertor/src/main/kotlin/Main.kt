import io.ktor.client.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import kotlin.math.roundToLong
import kotlin.system.exitProcess

@Serializable
data class GBP (val GBP: Float)

@Serializable
data class CurrencyData (val data: GBP)

object Coins {
    const val c500 = 500;
    const val c100 = 100;
    const val c50 = 50;
    const val c10 = 10;
    const val c5 = 5;
    const val c1 = 1;
}

object Notes {
    const val n10000 = 10000;
    const val n1000 = 1000;
    const val n5000 = 5000;
}

suspend fun main() {
    val wallet = arrayOf(
        Notes.n10000 * 5,
        Notes.n1000 * 5,
        Notes.n5000 * 1,
        Coins.c500 * 1,
        Coins.c100 * 4,
        Coins.c50 * 5,
        Coins.c10 * 8,
        Coins.c5 * 7,
        Coins.c1 * 14
    )

    val key: String = System.getenv("FREE_CURRENCY_API_KEY") ?: "";
    val url = "https://api.freecurrencyapi.com/v1/latest?apikey=${key}&currencies=GBP&base_currency=JPY"

    if (key.isEmpty()) {
        println("Please set you Free Currency API Key");
        exitProcess(1);
    }

    val client = HttpClient();
    val response: HttpResponse = client.get(url);
    val data = Json.decodeFromString<CurrencyData>(response.bodyAsText())
    val jpyToGbp = data.data.GBP;
    client.close();

    val total = wallet.reduce {a, e -> a + e}
    val totalGbp = total * jpyToGbp
    val totalGbp2dec = (totalGbp * 100f).roundToLong() / 100f

    println("You have ¥$total , which is £${totalGbp2dec}");

}