// swift-tools-version:5.1
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "AlibabaCloudFCUtils",
    products: [
        .library(
            name: "AlibabaCloudFCUtils",
            targets: ["AlibabaCloudFCUtils"])
    ],
    dependencies: [
        .package(url: "https://github.com/aliyun/tea-swift.git", from: "0.3.0"),
        .package(url: "https://github.com/krzyzanowskim/CryptoSwift.git", from: "1.3.0")
    ],
    targets: [
        .target(
            name: "AlibabaCloudFCUtils",
            dependencies: ["CryptoSwift", "Tea"]),
        .testTarget(
            name: "AlibabaCloudFCUtilsTests",
            dependencies: ["AlibabaCloudFCUtils", "Tea", "CryptoSwift"]),
    ]
)
