<template>
  <div class="h-[500px] rounded-lg">
    <l-map 
      ref="map" 
      v-model:zoom="zoom" 
      :center="center" 
      :use-global-leaflet="false" 
      :options="mapOptions" 
      style="height: 100%; width: 100%;"
      :max-bounds="maxBounds"
    >
        <l-tile-layer 
          :url="url" 
          :attribution="attribution"
        />
        <l-marker 
          v-for="event in events" 
          :key="event.id" 
          :lat-lng="[event.lat, event.lng]" 
          :icon="getIcon(event.id === hoveredEventId || event.id === selectedEventId)" 
          @click="onMarkerClick(event.id)"
          @mouseover="hoveredEventId = event.id" 
          @mouseout="hoveredEventId = null"
        />
    </l-map>
  </div>
</template>

<script setup>
import { LMap, LTileLayer, LMarker } from '@vue-leaflet/vue-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import { ref } from 'vue';

const greenIcon = new L.Icon({
  iconUrl: 'https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-green.png',
  shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-shadow.png',
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  shadowSize: [41, 41],
});

const redIcon = new L.Icon({
  iconUrl: 'https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-red.png',
  shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-shadow.png',
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  shadowSize: [41, 41],
});

const props = defineProps({
  events: { type: Array, required: true },
  selectedEventId: { type: Number, default: null },
});

const emit = defineEmits(['update:selected']);
const hoveredEventId = ref(null);

const center = ref([55.75, 37.62]);  // Центр России
const zoom = ref(3);
const mapOptions = ref({ attributionControl: false });

const url = 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png';
const attribution = '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors';

// Ограничиваем карту только территорией России
const maxBounds = [
  [41.1850968, 19.6389],  // Южный угол карты
  [82.0586232, 180]       // Северный угол карты
];

const getIcon = (isHighlighted) => isHighlighted ? redIcon : greenIcon;

// Click handler for actual selection of event
const onMarkerClick = (id) => {
  emit('update:selected', id);
  
  // Scroll to the event card if it exists
  const eventElement = document.getElementById(`event-${id}`);
  if (eventElement) {
    eventElement.scrollIntoView({ behavior: 'smooth' });
  }
};
</script>
