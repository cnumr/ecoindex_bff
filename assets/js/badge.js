(function(){
    const baseUrl = "{{url}}";
    const url = window.location.href;
    const badge = document.getElementById("ecoindex-badge");
    const a = document.createElement("a");
    const img = document.createElement("img"); 

    a.href = `${baseUrl}/redirect/?url=${url}`; 
    a.target = "_blank";
    img.src = `${baseUrl}/?badge=true&url=${url}`;
    a.appendChild(img);
    badge.appendChild(a);
})();

