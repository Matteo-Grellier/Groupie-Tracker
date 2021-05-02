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
  { lat: 47.21257651938386, lng: -1.5545689756191343}, 
  { lat:  47.25010637451308, lng: -1.520823648724015 }, 
  { lat:  47.22434944289329, lng: -1.62313012459583},
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





