<template>
  <!--Modal-->
  <div class="modal fixed w-full h-full top-0 left-0 flex items-center justify-center">
    <div class="modal-overlay absolute w-full h-full bg-gray-900 opacity-50"></div>
    
    <div class="modal-container bg-white w-11/12 md:max-w-md mx-auto rounded shadow-lg z-50 overflow-y-auto">
      
      <!-- Add margin if you want to see some of the overlay behind the modal-->
      <div class="modal-content py-4 text-left px-6">
        <!--Title-->
        <div class="flex justify-between items-center pb-3">
          <h1 class="text-xl font-bold text-black">New Concept</h1>
          <div class="modal-close cursor-pointer z-50">
            <svg class="fill-current text-black" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 18 18">
              <path d="M14.53 4.53l-1.06-1.06L9 7.94 4.53 3.47 3.47 4.53 7.94 9l-4.47 4.47 1.06 1.06L9 10.06l4.47 4.47 1.06-1.06L10.06 9z"></path>
            </svg>
          </div>
        </div>

        <!--Body-->
        <input />

        <!--Footer-->
        <div class="flex justify-end pt-2">
          <button class="px-4 primary-color py-2 rounded text-white shadow hover:shadow-lg font-medium capitalize">CREATE</button>
        </div>
        
      </div>
    </div>
  </div> 
</template>

<script>
export default {
  name: 'NewConceptModal',
}
</script>

<style scoped>

</style>
