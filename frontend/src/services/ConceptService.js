import { baseUrl } from './core.js'

export async function createConcept(title, content, prereqIds) {
  const req = {
    title,
    content,
    prerequisites: prereqIds
  }

  await fetch(`${baseUrl}/concept/new`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(req)
  })
}

export async function getAllConcepts() {
      // Go and retrieve the concepts
  const rawResponse = await fetch(`${baseUrl}/concept/all`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }}) 

  return await rawResponse.json()

}

export async function updateConcept(id, content) {
  // Run the update
  await fetch(`${baseUrl}/concept/update/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({content: content})
  })  
}

export async function deleteConcept(id) {
  await fetch(`${baseUrl}/concept/delete/${id}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    }})
}

export async function searchForConceptByTitle(query) {
  const rawResponse = await fetch(`${baseUrl}/concept/search?` + new URLSearchParams({query: query}).toString(), {
    method: 'GET'})
  return await rawResponse.json()
}
