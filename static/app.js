console.log("hi")

const form = document.forms[0]
form.addEventListener("submit", async (e)=>{
    e.preventDefault()
    const reqs = parseInt(form.reqs.value)
    const con = parseInt(form.con.value)
    const url = form.url.value
    try{
        const req = await fetch("/start",{
            method:"POST",
            headers:{
                "Content-Type":"application/json"
            },
            body:JSON.stringify({url, con, req:reqs})
        })
    }catch(err){
        console.log(err)
    }
})