app.controller('homeController', function ($scope, $http) {
    $scope.submit = function (data) {
        //Invoke CC address
        $scope.postparams = {
            "jsonrpc": "2.0",
            "method": "invoke",
            "params": {
                "type": 1,
                "chaincodeID": {
                    "name": "string"
                },
                "ctorMsg": {
                    "function": "string",
                    "args": [
                        "string"
                    ]
                },
                "secureContext": "string"
            },
            "id": 0
        }
    
}
$scope.goToView = function(){

    window.location = '#!/viewAsset'
}


});