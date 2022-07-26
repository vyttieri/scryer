<script setup>
	import { storeToRefs } from 'pinia'

	import { useDeviceStore } from '../stores/device'

	const { devices } =  storeToRefs(useDeviceStore())
</script>

<template>
	<GMapMap
		:center="center"
		:zoom="7"
		map-type-id="terrain"
		style="width: 100vw; height: 100%"
	>
		<GMapMarker
			:key="'device.device_id' + '-marker'"
			:position="device.latest_accurate_device_point.device_point_detail.lat_lng"
			v-for="device in devices"
		/>
	</GMapMap>

</template>

<style scoped>
</style>

<script>
export default {
	name: 'Map',
	data() {
		return {
			center: { lat: 51, lng: 7 },
		}
	}
}
</script>
