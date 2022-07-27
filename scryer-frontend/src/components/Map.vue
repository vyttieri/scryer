<script setup>
	import { storeToRefs } from 'pinia'

	import { useDeviceStore } from '../stores/device'
	import { usePreferencesStore } from '../stores/preferences'

	const { devices, visibleDevices } =  storeToRefs(useDeviceStore())
	const { center } = storeToRefs(usePreferencesStore())

</script>

<!-- TODO: Centering seems wrong -->
<template>
	<GMapMap
		:center="center"
		:zoom="6"
	>
		<GMapMarker
			:key="'device.device_id' + '-marker'"
			:position="device.latest_accurate_device_point.device_point_detail.lat_lng"
			v-for="device in visibleDevices"
		/>
	</GMapMap>
</template>
