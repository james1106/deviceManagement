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
    $scope.send = function () {
        $scope.deploySpec = {
            "jsonrpc": "2.0",
            "method": "deploy",
            "params": {
                "type": 1,
                "chaincodeID": {
                    "path": "https://github.com/nik0405/deviceManagement/chaincode"
                },
                "ctorMsg": {
                    "function": "init",
                    "args": [
                        JSON.stringify({ id: 'GENEMP001', name: 'SAM', assets: [] }),
                        JSON.stringify({ id: 'GENEMP002', name: 'RON', assets: [] }),
                        JSON.stringify({ id: 'GENEMP003', name: 'TOM', assets: [] }),
                        JSON.stringify({ make: 'Dell', type: 'Laptop', model: 'Latitude 5470', serialNo: 'GEN1LAP001', empId: '' }),
                        JSON.stringify({ make: 'Apple', type: 'Laptop', model: 'Macbook Air', serialNo: 'GEN1LAP002', empId: '' }),
                        JSON.stringify({ make: 'Apple', type: 'Laptop', model: 'Macbook Pro', serialNo: 'GEN1LAP003', empId: '' })
                    ]

                },
                "secureContext": "user_type1_0"
            },
            "id": 0
        }
        $http.post('https://d8e451794079457c973c666715f41a74-vp0.us.blockchain.ibm.com:5003/chaincode', $scope.deploySpec).then(function (response) {
            if (response.status === 200) {
               }
            return response;
        }, function (response) {
            return response.status
        });


    }
    $scope.goToView = function () {

        window.location = '#!/viewAsset'
    }


});