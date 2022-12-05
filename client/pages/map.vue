<template>
  <Base/>
</template>

<script>
export default {
  name: 'map',
  head() {
    return {
      title: 'Expedition Diaries countries map',
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: 'countries map'
        }
      ]
    }
  },
}
</script>

<style src="vue-svg-map/dist/index.css"></style>
