<template>
  <div class="relative items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link 
      href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css"
      rel="stylesheet" 
      class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container custom-container">
      <Header/>
      <div class="mt-2 bg-white overflow-hidden shadow sm:rounded-lg p-6">
        <h2 class="text-2xl leading-7 font-semibold">
          The expedition diaries' magical map!
        </h2>
        <p class="mt-3 text-gray-600">
        Select your favorite country to check it's expedition diaries.
        </p>
        <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
          <Map/>
        </div>
      </div>
     <Footer/> 
    </div>
  </div>
</template>

<script>
import Map from './Map.vue'
import Header from './Header.vue'
import Footer from './Footer.vue'

export default {
  name: 'Base',
  components: {
    Map,
    Header,
    Footer
  },
}
</script>
<style>
.custom-container {
  width: 50rem;
}
</style>