app.controller('viewAssetController', function ($scope, $http) {
   $scope.url = 'https://d8e451794079457c973c666715f41a74-vp0.us.blockchain.ibm.com:5003/';
   $scope.ccId ='93ea4891e341b6125b11629c3bf22f0c4b5855e395379f677600fd1c4ca6cd48948c599fe808cfb7cbe2d5ca58cb11734e403d4a9f85443d944bad9639d57571';
  
    $scope.getAssets = function () {
        $scope.postparams = {
            "jsonrpc": "2.0",
            "method": "query",
            "params": {
                "type": 1,
                "chaincodeID": {
                    "name": $scope.ccId
                },
                "ctorMsg": {
                    "function": "query",
                    "args": [
                        "assets"
                    ]
                },
                "secureContext": "user_type1_0"
            },
            "id": 0
        }
           $http.post($scope.url + '/chaincode', $scope.postparams).then(function (response) {
                $scope.assetContainer = JSON.parse(response.data.result.message);
                return response
            }, function (response) {
                return response.status
            })
    
}
$scope.getAssets(); 

});