// carte interactive

let map;

function initMap() {
  map = new google.maps.Map(document.getElementById("map"), {
    center: { lat: 47.20565350041161, lng: -1.5389924790850191},
    zoom: 8,
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
];


