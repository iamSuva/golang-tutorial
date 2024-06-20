const express=require("express");
const app=express();
app.use(express.json())
app.use(express.urlencoded({extended:true}))
app.get("/get",(req,res)=>{
    return res.status(200).send({message:"Welcome to nodejs server"})
})

app.post("/post",(req,res)=>{
    const data=req.body;
    console.log(data);
    return res.send(data);
})
app.post("/postform",(req,res)=>{
    const data=req.body;
    return res.send(JSON.stringify(data));
})
app.listen(4000,()=>{
    console.log("listening on port 4000");
})