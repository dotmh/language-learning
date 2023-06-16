//
//  main.swift
//  currency
//
//  Created by Martin Haynes on 16/06/2023.
//

import Foundation

let url = URL(string: "https://learn.api.dotmh.dev")!

var request = URLRequest(url: url);
request.setValue("application/json", forHTTPHeaderField: "Content-Type")

print("Hello, World!")

