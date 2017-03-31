app.controller('enrollController', function ($scope, $http, $q) {
    $scope.url = [
        "https://0c3794f948bc4cf595d83111520fb20c-vp0.us.blockchain.ibm.com:5004",
        "https://0c3794f948bc4cf595d83111520fb20c-vp1.us.blockchain.ibm.com:5004",
        "https://0c3794f948bc4cf595d83111520fb20c-vp2.us.blockchain.ibm.com:5004",
        "https://0c3794f948bc4cf595d83111520fb20c-vp3.us.blockchain.ibm.com:5004"
    ]

    $scope.users = [{
            "enrollId": "user_type1_0",
            "enrollSecret": "e75d391289"
        },
        {
            "enrollId": "user_type1_1",
            "enrollSecret": "cfd7124739"
        },
        {
            "enrollId": "user_type1_2",
            "enrollSecret": "1e7dd449bc"
        },
        {
            "enrollId": "user_type1_3",
            "enrollSecret": "d559778d79"
        }
    ]

    var defer = $q.defer();

    $scope.enroll = function (data) {

        $q.all([
            $http.post($scope.url[0] + '/registrar', data).then(function (response) {
                $scope.loggedInUser = response.status;
                return response
            }, function (response) {
                return response.status
            }),
            $http.post($scope.url[1] + '/registrar', $scope.users[1]).then(function (response) {
                return response;
            }, function (response) {
                alert('VP1 NON RESPONSIVE');
                return response.status
            }),
            $http.post($scope.url[2] + '/registrar', $scope.users[2]).then(function (response) {
                return response;
            }, function (response) {
                alert('VP2 NON RESPONSIVE');
                return response.status
            }),
            $http.post($scope.url[3] + '/registrar', $scope.users[3]).then(function (response) {
                return response;
            }, function (response) {
                alert('VP3 NON RESPONSIVE');
                return response.status
            })
        ]).then(function (array) {
            if ($scope.loggedInUser === 200) {
                //Redirect to home
                window.location = '#/home'
            }
        });
    }
});

