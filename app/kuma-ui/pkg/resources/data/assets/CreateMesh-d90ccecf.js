import{W as l}from"./kongponents.es-8abed680.js";import{P as u}from"./production-0f1ffdb6.js";import{u as _}from"./store-0511bcbf.js";import{O as m,a as f,b as h}from"./OnboardingPage-8b536a7f.js";import{d as b,r as g,c as v,a as y,w as a,o as x,e as s,f as e,g as n,t as S,u as o,p as A,j as C}from"./runtime-dom.esm-bundler-a6f4ece5.js";import{_ as N}from"./_plugin-vue_export-helper-c27b6911.js";const r=t=>(A("data-v-fdce75bc"),t=t(),C(),t),k={class:"mb-4 text-center"},B=r(()=>n("i",null,"default",-1)),D=r(()=>n("p",{class:"mt-4 text-center"},`
        This mesh is empty. Next, you add services and their data plane proxies.
      `,-1)),M=b({__name:"CreateMesh",setup(t){const c=[{label:"Name",key:"name"},{label:"Services",key:"servicesAmount"},{label:"DPPs",key:"dppsAmount"}],i=_(),d=g({total:1,data:[{name:"default",servicesAmount:0,dppsAmount:0}]}),p=v(()=>i.getters["config/getMulticlusterStatus"]?"onboarding-multi-zone":"onboarding-configuration-types");return(O,P)=>(x(),y(m,null,{header:a(()=>[s(f,null,{title:a(()=>[e(`
          Create the mesh
        `)]),_:1})]),content:a(()=>[n("p",k,[e(`
        When you install, `+S(o(u))+" creates a ",1),B,e(` mesh, but you can add as many meshes as you need.
      `)]),e(),s(o(l),{class:"table",fetcher:()=>d.value,headers:c,"disable-pagination":""},null,8,["fetcher"]),e(),D]),navigation:a(()=>[s(h,{"next-step":"onboarding-add-services","previous-step":o(p)},null,8,["previous-step"])]),_:1}))}});const H=N(M,[["__scopeId","data-v-fdce75bc"]]);export{H as default};
