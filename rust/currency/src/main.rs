async fn getJpyToGbp() {
    let body = reqwest::get("https://www.dotmh.io")
        .await?
        .text()
        .await?;
}

fn main() {


    println!("Body {:?}", body);
}
