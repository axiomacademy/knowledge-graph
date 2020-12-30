<template>
  <h1 class="text-3xl font-medium pt-2 text-black">{{ title }}</h1>
  <div class="flex flex-row px-8 items-center">
    <h1 class="text-xl font-medium mr-auto text-gray-900">Concept Content</h1>
    <button @click="onSave" class="inline-block px-6 py-2 text-xs font-medium leading-6 text-center text-white uppercase transition primary-color rounded shadow ripple hover:shadow-lg focus:outline-none">
      Save
    </button>
  </div>
  <div class="flex flex-row mx-8 my-4 border border-gray-200 border-solid max-h-screen rounded">
    <textarea id="raw" :value="content" @input="update" class="flex-1 resize-none bg-gray-100 p-8"></textarea>
    <div id="processed" v-html="compiledMarkdown" class="flex-1 p-8 text-left markdown-body"></div>
  </div> 
</template>

<script>
import * as marked from 'marked'

export default {
  name: 'MarkdownEditor',
  props: {
    title: String,
    content: String,
  },
  data () {
    return {
    }
  },
  computed: {
    compiledMarkdown: function() {
      return marked(this.content);
    }
  },
  methods: {
    update(e) { 
      this.$emit("update:content", e.target.value);
    },
    onSave() {
      console.log("Saving...")
      this.$emit("save", this.input)
    },
  }
}
</script>

<style>
#raw {
  font-family: 'Roboto Mono', monospace;
}
</style>
