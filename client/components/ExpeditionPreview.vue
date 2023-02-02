<template>
  <div>
    <a :href="`/travels/`+slug" class="underline text-blue-600 hover:text-blue-800 visited:text-purple-600">
      <h2 class="font-bold text-xl">{{title}}</h2>
    </a>
    <div v-if="description" class="whitespace-pre-line"><p>{{ description }}</p></div> 
    <br>

    <h2 class="font-bold">📅 Duration</h2>
    <div>
      <p>From {{startedAt? startedAt : "?"}} A.D. to {{endedAt? endedAt : "?"}} A.D.</p>
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

    <div class="flex space-x-2 justify-end">
      <a :href="`/travels/`+slug" target="_blank">
        <button type="button" class="inline-block px-6 py-2.5 bg-blue-500 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Read More</button>
      </a>
    </div>
  </div>
</template>

<script>

export default {
  props: ['title', 'description', 'route', 'startedAt', 'endedAt', 'explorers', 'slug']
}
</script>
