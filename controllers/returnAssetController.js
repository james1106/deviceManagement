app.controller('returnAssetController', function ($scope, $http, $q ,$window) {
    $scope.url = 'https://d8e451794079457c973c666715f41a74-vp0.us.blockchain.ibm.com:5003/';
    $scope.ccId ='93ea4891e341b6125b11629c3bf22f0c4b5855e395379f677600fd1c4ca6cd48948c599fe808cfb7cbe2d5ca58cb11734e403d4a9f85443d944bad9639d57571';

    $scope.queryEmpSpec = {
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
                    "employees"
                ]
            },
            "secureContext": "user_type1_0"
        },
        "id": 0
    }
    $scope.queryAssetSpec = {
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
    var defer = $q.defer();

    $scope.getData = function () {

        $q.all([
            $http.post($scope.url + '/chaincode', $scope.queryEmpSpec).then(function (response) {
                $scope.empList = JSON.parse(response.data.result.message);
                console.log($scope.empList)
                return response
            }, function (response) {
                return response.status
            }),
            $http.post($scope.url + '/chaincode', $scope.queryAssetSpec).then(function (response) {
                $scope.assetList = JSON.parse(response.data.result.message);
                console.log($scope.assetList)
                return response;
            }, function (response) {
                return response.status
            })
        ]).then(function (array) {

        });

    }

    $scope.getData();
    $scope.$watch('employees', function (newValue, oldValue) {
        if (newValue) {
            $scope.correspondingEmployeAssets = newValue.assets;
        }
    });
  
     $scope.returnAsset = function (selEmp,selAsset) {
        $scope.assignAssetSpecs = {
            "jsonrpc": "2.0",
            "method": "invoke",
            "params": {
                "type": 1,
                "chaincodeID": {
                    "name": $scope.ccId
                },
                "ctorMsg": {
                    "function": "returnAsset",
                    "args": [
                        JSON.stringify(selEmp), JSON.stringify(selAsset)
                    ]
                },
                "secureContext": "user_type1_0"
            },
            "id": 0
        }
        console.log($scope.assignAssetSpecs)
        $http.post($scope.url + '/chaincode', $scope.assignAssetSpecs).then(function (response) {
            if (response.status === 200) {
                alert("Asset returned")
                $window.location.reload();
            }
            return response
        }, function (response) {
            return response.status
        })
    }
});