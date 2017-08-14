let vm = new Vue({
            el: '#wrapper',   
            data : {
                currencies : [ ]
            },
            mounted: function () {

                createWebSocket()

            }
        })

        function createWebSocket() {

            let wsuri = "ws://127.0.0.1:8080/ws"; 
            let ws = new WebSocket(wsuri);

            ws.onopen = function() {
                console.log("Connected to " + wsuri);
            }

            ws.onclose = function(e) {
               console.log("Connection closed (" + e.code + ")");
               
               // TRY TO RECONNECT EVERY 5 SEC
               setTimeout( function() {
                   createWebSocket()
               }, 5000);

            }

            ws.onmessage = function(e) {
                let result = JSON.parse(e.data)

                vm.currencies = result.Result.sort( function(a, b) { return b.BaseVolume - a.BaseVolume } );

                console.log(vm.currencies)
            }
        }