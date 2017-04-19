app.controller('addAssetController', function ($scope, $http,$window) {
    $scope.url = 'https://d8e451794079457c973c666715f41a74-vp0.us.blockchain.ibm.com:5003/';
    $scope.ccId ='93ea4891e341b6125b11629c3bf22f0c4b5855e395379f677600fd1c4ca6cd48948c599fe808cfb7cbe2d5ca58cb11734e403d4a9f85443d944bad9639d57571';
      $scope.added=false;
    $scope.addDevice = function (device) {
        console.log(device)
        var data = {
            make: device.make,
            model: device.model,
            type: device.type,
            serialNo: device.serial
        }
        console.log(data)
        $scope.postparams = {
            "jsonrpc": "2.0",
            "method": "invoke",
            "params": {
                "type": 1,
                "chaincodeID": {
                    "name": $scope.ccId
                },
                "ctorMsg": {
                    "function": "addAsset",
                    "args": [
                        JSON.stringify(data)
                    ]
                },
                "secureContext": "user_type1_0"
            },
            "id": 0
        }

        $http.post($scope.url + 'chaincode', $scope.postparams).then(function (response) {
            if (response.status === 200) {
               $scope.added=true;
               $scope.newDevice =null;
                $window.location.reload();
            }
            return response;
        }, function (response) {
              $scope.added=false;
            return response.status
        });

    }
});