var app = angular.module("app", ["ngRoute"]);
app.config(function ($routeProvider) {
    $routeProvider
        .when("/", {
            templateUrl: "views/enroll.html",
            controller: "enrollController"
        })
        .when("/home", {
            templateUrl: "views/home.html",
            controller: "homeController"
        })
        .when("/viewAsset", {
            templateUrl: "views/viewAsset.html",
            controller: "viewAssetController"
        })
         .when("/addAsset", {
            templateUrl: "views/addAsset.html",
            controller: "addAssetController"
        })
         .when("/assignAsset", {
            templateUrl: "views/assignAsset.html",
            controller: "assignAssetController"
        })
         .when("/returnAsset", {
            templateUrl: "views/returnAsset.html",
            controller: "returnAssetController"
        });
});