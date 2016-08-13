window.like = (type, id) => $.post(`/api/${type}/like/` + id)
window.unlike = (type, id) => $.post(`/api/${type}/unlike/` + id)