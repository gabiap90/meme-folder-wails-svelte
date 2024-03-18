import { writable } from "svelte/store";

export const images = writable([])
export const selectedTags = writable([])
export const allTags = writable([])

export function updateImages(res) {
    images.set(res)
    allTags.set([...new Set(res.map((image) => image.Tags).flat())]);
}