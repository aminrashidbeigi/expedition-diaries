<template>
   <div class="relative items-top justify-center min-h-screen bg-gray-100 sm:items-center sm:pt-0">
    <link 
      href="https://cdn.jsdelivr.net/npm/tailwindcss@2.1.2/dist/tailwind.min.css" 
      rel="stylesheet"
      class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
    <div class="max-w-4xl mx-auto sm:px-6 lg:px-8 container">
      <Header/>
      <div class="mt-8 bg-white overflow-hidden shadow sm:rounded-lg p-6">
        <h2 class="text-2xl leading-7 font-bold">
          Add an expedition to Expedition Diaries
        </h2>
        <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
          Please add details of an expedition within
            <a 
              href="https://forms.gle/JWDhVp2uoiqAfHCBA" 
              class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
               this form
            </a>.
          <br>
          We will let you know whenever the expedition added to the database :)
        </div>
      </div>
      <Footer/>
    </div>
  </div>
</template>

<script>
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'

export default {
  components: {
    Header,
    Footer,
  },
}
</script>
