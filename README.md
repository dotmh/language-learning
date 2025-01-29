# Language Learning Scenario

## About

I enjoy learning new languages in pursuit of that goal I came up with a basic scenario that I could use to explore new languages and frameworks.

The idea of the scenario is that it is simple enough to not require in depth knowledge of the language so making it easier to learn the language. It also can cover some common tasks that I would need to do as a front end developer.

## The Scenario

You have returned from a 2 week trip in Japan. You didn't spend all the money that you convert to the Japanese Yen (¥) to take with you. You there for have a wallet fall of Japanese currency. The task is to covert the ¥ amount back to British Pound Sterling (£).

### The code should

1. take in the amount of each coin or note that you have.
2. show a summary of the wallet totals i.e. how much value of each coin or note you have
3. get the total in ¥
4. get the current exchange rate from [Free Currency API](https://freecurrencyapi.com) for ¥ to £.
5. display the total in ¥ and in £ to 2 decimal places.

### Wallet Contents

#### Notes

- 5 x ¥10,000 notes
- 1 x ¥5000 note
- 5 \* ¥1000 notes

#### Coins

- 1 x ¥500 coin
- 4 x ¥100 coins
- 5 x ¥50 coins
- 8 x ¥10 coins
- 7 x ¥5 coins
- 14 x ¥1 coins

## Scope

### Reducing Scope

This is already a simple task, but with some languages or situations making HTTP requests can be very complex.

You can therefore reduce scope by removing the part around getting a live conversion rate for ¥ to £.
You then replace the live exchange rate with a fixed value.

### Extending Scope

There are number of ways that scope can be expand

#### User Input

As an example app, having the wallet (the amount of coins and notes that you have) hardcoded is fine
but in the real world it would be better to allow the user to input that data. so you could add

- allow the program to loop through each note and coin and ask the user how much they have.

#### Graphical User Interface

The next cool step might be to add a Graphical User Interface (GUI). Real Users prefer a GUI as they find Command Line Interface (CLI) apps much more intimidating. In this case you might add

- show the user graphical the total in ¥ and in £

you could also combine this with the first scope extension to allow the user to enter the how much of
each coin or note they have inside your GUI.

- allow the user to add how much of each coin or note they have using the GUI.

#### Expanding the Currencies

The Current app is very much focused on a single purpose that off converting ¥ to £ but to expand it
in to a much more useful application might be to add in extra currencies. This could take the form of
adding extra currencies to convert from , extra currencies to convert too or both. This is the least
interesting from a coding point of view because it will be more about Data. You would need to get
all the different coins and notes used by each currency.

The biggest logic change will be handling the sub unit of currencies used by some currencies such as
Cents (¢) on the Dollar ($) or pence (p) on the Pound (£). Although again this is pretty simple as almost all
currencies have a 100 sub units of currency to the main.

With all of these scope extensions added together you would have a full app that could be useful for
any user.

- Allow the user to change which currency they want to convert to
- Allow the user to change which currency the user wants to convert from
- Show the correct Notes and Coins for the currency the user has chosen to convert from.

## API

You can get the above data, as well as Japanese coin and note denominations from an API endpoint.

Make a [HTTP GET](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET)
request to <https://learn.api.dotmh.dev>

[![Deno JS](https://img.shields.io/badge/deno%20js-000000?style=for-the-badge&logo=deno&logoColor=white)](/support/api/api.ts)

## Implementations

Please note these implementations (with the exemption of Typescript) are all created while learning and so may not reflect best practise. These are not intended to be definitive examples of each language.

[![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)](/typescript/node/index.ts)
[![Kotlin](https://img.shields.io/badge/kotlin-%237F52FF.svg?style=for-the-badge&logo=kotlin&logoColor=white)](/kotlin/currency-convertor/src/main/kotlin/Main.kt)
[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](/go/currency.go)
[![C#](https://img.shields.io/badge/c%23-%23239120.svg?style=for-the-badge&logo=csharp&logoColor=white)](cs/curreny-calculator/)

### Notes for running

#### Environment Variables

All examples expect the API key for [Free Currency API](https://freecurrencyapi.com) to be stored in the Environment variable `FREE_CURRENCY_API_KEY`
