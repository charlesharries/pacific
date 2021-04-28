export default function useCsrf(): string {
  const $meta: HTMLMetaElement | null = document.querySelector('meta[name="csrf_token"]');

  return $meta ? $meta.content : '';
}
