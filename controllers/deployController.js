app.controller('deploymentController', function ($scope, $http) {
    $scope.submit = function (data) {
        //Initializing initial employees
        $scope.employees = [{
            id: 1,
            name: "SAM",
            assets: [{make:"Dell",type: "Laptop",model: "Latitude 5470",serialNo: "GEN1LAP001"}]
        },
        {
            id: 2,
            name: "RON",
            assets: [{make:"Apple",type: "Laptop",model: "Macbook Air",serialNo: "GEN1LAP002"}]
        }, {
            id: 3,
            name: "TOM",
            assets: [{make:"Apple",type: "Laptop",model: "Macbook Pro",serialNo: "GEN1LAP003"}]
        }, {
            id: 4,
            name: "ERIC",
            assets: []
        }, {
            id: 5,
            name: "ANTONY",
            assets: []
        }];
        //Initializing initial assets
        $scope.assets = [
            //Laptops
            {
                make: "DELL",
                type: "Laptop",
                model: "Latitude 5470",
                serialNo: "GENLAP001"
            },
            {
                make: "Apple",
                type: "Laptop",
                model: "Macbook Air",
                serialNo: "GENLAP002"
            },
            {
                make: "Apple",
                type: "Laptop",
                model: "Macbook Pro",
                serialNo: "GENLAP003"
            },
            //Mobiles
            {
                make: "Apple",
                type: "Mobile",
                model: "iPhone 6",
                serialNo: "GENMOB001"
            },
            {
                make: "Microsoft",
                type: "Mobile",
                model: "Lumia 520",
                serialNo: "GENMOB002"
            },
            {
                make: "Samsung",
                type: "Mobile",
                model: "Galaxy S7",
                serialNo: "GENMOB003"
            },
            //Desktops
            { 
                make: "Apple",
                type: "Desktop",
                model: "Mac Mini",
                serialNo: "GENDESKTOP001"
            },
            { 
                make: "Dell",
                type: "Desktop",
                model: "Vostro",
                serialNo: "GENDESKTOP002"
            },
            { 
                make: "Dell",
                type: "Desktop",
                model: "Vostro",
                serialNo: "GENDESKTOP003"
            },
            //Tablets
             { 
                make: "Apple",
                type: "Tablet",
                model: "iPad Mini",
                serialNo: "GENTAB001"
            },
             { 
                make: "Apple",
                type: "Tablet",
                model: "iPad",
                serialNo: "GENTAB002"
            },
             { 
                make: "Samsung",
                type: "Tablet",
                model: "Galaxy Mini",
                serialNo: "GENTAB003"
            },
             { 
                make: "Samsung",
                type: "Tablet",
                model: "Galaxy",
                serialNo: "GENTAB004"
            },
            //Mouse
            { 
                make: "Dell",
                type: "Mouse",
                model: "Optical Wireless",
                serialNo: "GENMOUSE001"
            },
              
              { 
                make: "Dell",
                type: "Mouse",
                model: "Optical Wireless",
                serialNo: "GENMOUSE002"
            },  
            { 
                make: "Dell",
                type: "Mouse",
                model: "Optical Wireless",
                serialNo: "GENMOUSE003"
            },
              { 
                make: "Dell",
                type: "Mouse",
                model: "Optical Wireless",
                serialNo: "GENMOUSE004"
            },
              { 
                make: "Dell",
                type: "Mouse",
                model: "Optical Wireless",
                serialNo: "GENMOUSE005"
            },
              { 
                make: "Dell",
                type: "Mouse",
                model: "Optical Wireless",
                serialNo: "GENMOUSE006"
            },
              { 
                make: "Dell",
                type: "Mouse",
                model: "Optical Wireless",
                serialNo: "GENMOUSE007"
            },
              { 
                make: "Dell",
                type: "Mouse",
                model: "Optical Wireless",
                serialNo: "GENMOUSE008"
            },
              { 
                make: "Apple",
                type: "Mouse",
                model: "Magic",
                serialNo: "GENMOUSE009"
            },
              { 
                make: "Apple",
                type: "Mouse",
                model: "Magic",
                serialNo: "GENMOUSE0010"
            },
        ];





        //Invoke CC address
        $scope.deploySpec = {
            "jsonrpc": "2.0",
            "method": "deploy",
            "params": {
                "type": 1,
                "chaincodeID": {
                    "path": "https://github.com/nik0405/GlobalPayments1.1"
                },
                "ctorMsg": {
                    "function": "init",
                    "args": ['{"id": "1","name": "SAM","assets": ["Dell","Laptop","Latitude 5470","GEN1LAP001"]}',
                        '{"id": "2","name": "RON","assets": ["Apple","Laptop","Macbook Air","GEN1LAP002"]}',
                        '{"id": "3","name": "TOM","assets": ["Apple","Laptop","Macbook Pro","GEN1LAP003"]}'
                    ]
                },
                "secureContext": "user_type1_0"
            },
            "id": 0
        }
        $http.post('https://f48a4804199c419aa3641b46e94e2dac-vp0.us.blockchain.ibm.com:5003/chaincode', $scope.deploySpec).then(function (response) {
            if (response.status === 200) {
                window.localStorage.setItem("chaincodeid", response.data.result.message)
                window.localStorage.setItem("custName", data.name)
                window.location = '#!/dashboard'
            }
            return response;
        }, function (response) {
            return response.status
        });

    }
});