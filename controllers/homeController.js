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
$scope.send = function(){
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
                    "args": [JSON.stringify({id: 1,name: 'SAM',assets: [{make:'Dell',type:'Laptop',model:'Latitude 5470',serialNo:'GEN1LAP001'}]}),
                        JSON.stringify({id: 2,name: 'RON',assets: [{make:'Apple',type:'Laptop',model:'Macbook Air',serialNo:'GEN1LAP002'}]}),
                        JSON.stringify({id: 3,name: 'TOM',assets: [{make:'Apple',type:'Laptop',model:'Macbook Pro',serialNo:'GEN1LAP003'}]})
                     ]
                },
                "secureContext": "user_type1_0"
            },
            "id": 0
        }
         $http.post('https://81b378259d304e18a7a593d4ebfc0590-vp0.us.blockchain.ibm.com:5001/chaincode', $scope.deploySpec).then(function (response) {
            if (response.status === 200) {
               console.log(response)
               
            }
            return response;
        }, function (response) {
            return response.status
        });


}    
$scope.goToView = function(){

    window.location = '#!/viewAsset'
}


});