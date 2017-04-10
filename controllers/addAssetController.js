app.controller('addAssetController', function ($scope, $http) {
//     $scope.submit = function (data) {
//         //Invoke CC address
//         $scope.postparams = {
//             "jsonrpc": "2.0",
//             "method": "invoke",
//             "params": {
//                 "type": 1,
//                 "chaincodeID": {
//                     "name": "string"
//                 },
//                 "ctorMsg": {
//                     "function": "string",
//                     "args": [
//                         "string"
//                     ]
//                 },
//                 "secureContext": "string"
//             },
//             "id": 0
//         }
    
// }
$scope.addDevice = function(device){
console.log(device)
 var data ={
            make:device.make,
            model:device.model,
            type:device.type,
            serialNo:device.serial
        }
    console.log(data)
 $scope.postparams = {
            "jsonrpc": "2.0",
            "method": "invoke",
            "params": {
                "type": 1,
                "chaincodeID": {
                    "name": "fb9525fa66b16bb59264572a9abd532023b2b29af93f9a11d8e05d2da3f9d55ddda430fac5349d91fb2621125d2f03eabe1724e0d46c3c4f454c46a30c61850e"
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
       
    $http.post('https://71af638568864f959f03a41d0938afbd-vp0.us.blockchain.ibm.com:5002/chaincode', $scope.postparams).then(function (response) {
            if (response.status === 200) {
               console.log(response)
               
            }
            return response;
        }, function (response) {
            return response.status
        });

}
});