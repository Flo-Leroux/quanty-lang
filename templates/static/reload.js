const source = new EventSource("/__sse");

source.addEventListener(
  "reload",
  function (e) {
    console.log("Reload", e);
    window.location.reload();
  },
  false
);

source.addEventListener(
  "open",
  function (e) {
    console.log("Connection was opened");
  },
  false
);

source.addEventListener(
  "error",
  function (e) {
    console.log("Conection error");
  },
  false
);
