// carte interactive

let map;

function initMap() {
  map = new google.maps.Map(document.getElementById("map"), {
    center: { lat: 47.20565350041161, lng: -1.5389924790850191},
    zoom: 8,
  });

  const geocoder = new google.maps.Geocoder();
  document.getElementById("submit").addEventListener("click", () => {
    geocodeAddress(geocoder, map);
  });

  // permet de marqué les markers
  const labels = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
 
  // permet de definir la position des markers
  const markers = locations.map((location, i) => {
    return new google.maps.Marker({
      position: location,
      label: labels[i % labels.length],
    });
  });

  // on ajoute un point qui groupe tout les markers
  new MarkerClusterer(map, markers, {
    imagePath:
      "https://developers.google.com/maps/documentation/javascript/examples/markerclusterer/m",
  });
}

// les coordonnées d'où sont placés les markers
const locations = [
  { lat: 43.94975057362789, lng: -120.54099111425059},
  { lat:  49.287506896615255, lng: -123.12012906125756 },
  { lat:  38.92646786709714, lng: -116.39867568020664}, 
  { lat:  39.76629434677865, lng: -105.81604926480364},
  { lat:  48.139999900290256, lng: 11.581659544308058},
  { lat:  50.08338374248348, lng: 14.439525082933736},
  { lat:  45.47691867286066, lng: 9.185122225349703},
];


// geocoding -> recherche une adresse sur la map google
function geocodeAddress(geocoder, resultsMap) {
  const address = document.getElementById("address").value;
  geocoder.geocode({ address: address }, (results, status) => {
    if (status === "OK") {
      resultsMap.setCenter(results[0].geometry.location);
      new google.maps.Marker({
        map: resultsMap,
        position: results[0].geometry.location,
      });
    } else {
      alert("Could not find adress: " + status);
    }
  });
}


//geolocalisation -> connaitre nos coordonnées

const successCallback = (position) => {
  console.log(position);
};

const errorCallback = (error) => {
  console.log(error);
};

navigator.geolocation.getCurrentPosition(successCallback, errorCallback, {
  enableHighAccuracy: true,
  timeout: 5000
});





