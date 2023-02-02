<template>
  <div class="mt-2 bg-white overflow-hidden shadow sm:rounded-lg p-6">
    <h1 class="text-2xl leading-7 font-bold">
      {{ title }}
    </h1>
    <div class="mt-4 pt-4 text-gray-800 border-t border-dashed">
        <div v-if="description" class="whitespace-pre-line"><p>{{ description }}</p></div> 
        <br>

        <h2 class="font-bold">📅 Duration</h2>
        <div>
          <p>From {{startedAt? startedAt : "?"}} to {{endedAt? endedAt : "?"}}</p>
        </div>
        <br>
      
        <h2 class="font-bold">🤠 Explorers</h2>
        <div v-for="traveler in explorers">
          <p v-if="traveler.Link">
            - <a :href="traveler.Link" target="_blank" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
              {{traveler.Name}}
            </a>
          </p>
          <p v-else>
            - {{traveler.Name}}
          </p>
          <div v-if="traveler.Image" class="relative mx-auto w-80">
            <img :src="$config.baseImageURL+traveler.Image" :alt="traveler.Name + ' image'" class="relative w-full h-full rounded-md shadow-md">
          </div>
        </div>
        <br>
    
        <div v-if="route">
          <h2 class="font-bold">📍 Route</h2>
          <div class="relative mx-auto w-80">
            <a :href="$config.baseImageURL+route" target="_blank">
              <img :src="$config.baseImageURL+route" :alt="title + ' expedition route'" class="relative w-full h-full rounded-md shadow-md">
            </a>
          </div>
          <br>
        </div>
    
        <h2 class="font-bold">🌎 Countries of this travel</h2>
        <ul>
          <li v-for="country in countries">
            - <a :href="`/countries/` + country.Code" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
              {{country.Name}}
            </a>
          </li> 
        </ul>
        <br>

        <h2 class="font-bold">📖 Resources</h2>
        <div v-for="resource in resources">
          <p class="font-meduim">
            <a :href="resource.Link" target="_blank" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
              {{resource.Title}}
            </a>
          </p>
          <div class="relative mx-auto w-60 h-80">
            <a :href="resource.Link" target="_blank">
              <img :src="$config.baseImageURL+resource.Image" :alt="resource.Title" class="relative w-full h-full rounded-md shadow-md">
            </a>
          </div>
          <br>
        </div>
    </div>
  </div>
</template>

<script>

export default {
  props: ['title', 'description', 'route', 'startedAt', 'endedAt', 'explorers', 'countries', 'resources']
}
</script>
