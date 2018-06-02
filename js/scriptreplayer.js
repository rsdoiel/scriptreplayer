console.log("Is this loading?");
(function (window, document) {
    'use strict';
    console.log("Hello World!");

    function run(element_id, typescript_url) {
        var elem = document.getElementById(element_id), 
            pause = false,
            term = new Terminal();


        console.log("DEBUG getting " + typescript_url + " from " + window.location);
        term.open(elem);
        term.write("DEBUG getting " + typescript_url + " from " + window.location + "\r\n");

        function nibble(src, count) {
            var s = src.slice(0, count);
            term.write(s);
            console.log("DEBUG nibble: ", s);
            s = src.slice(count);
            console.log("DEBUG typeof(s): ", typeof(s));
            return s;
        }

        function dribble(buf, timing) {
            var cnt = 0,
                delay = 0.0,
                t = {},
                head = '',
                tail = '';

            nibble = function() {
                term.write(head);
                console.log("nibble: ", head);
                dribble(tail, timing);
            };

            t = timing.shift();
            if (t) {
                cnt = t.c;
                delay = t.t;
                head = buf.slice(0, cnt);
                tail = buf.slice(cnt);
                if (delay < 0.001) {
                    nibble();
                } else {
                    setTimeout(nibble, delay * 800);
                }
            }
        }

        var oReq = new XMLHttpRequest();
        oReq.addEventListener("load", function() {
            var src = this.responseText,
                o = {};
            console.log("DEBUG this.status?", this.status);
            console.log("DEBUG getting src", src); 
            o = JSON.parse(src);
            console.log("DEBUG calling dribble()"); 
            dribble(o.typescript, o.timing);
        });
        oReq.open("GET", typescript_url + "?time=" + (new Date()).getTime());
        oReq.send(null);
        console.log("DEBUG all done!");
    }

    // Export our object
    window.scriptreplayer = {};
    window.scriptreplayer.run = run;
    console.log("DEBUG export done!");
}(window, document));
