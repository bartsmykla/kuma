import{e as u}from"./index-fce48c05.js";import{d as _,a as n,o as f,b as y,w as i,e as c,l as T,ad as x,f as C,r as m,ae as h,m as B,t as b,_ as v}from"./index-a842a987.js";const g={class:"visually-hidden"},w={inheritAttrs:!1},S=_({...w,__name:"CopyButton",props:{text:{default:""},getText:{type:[Function,null],default:null},copyText:{default:"Copy"},tooltipSuccessText:{default:"Copied code!"},tooltipFailText:{default:"Failed to copy!"},hasBorder:{type:Boolean,default:!1},hideTitle:{type:Boolean,default:!1},iconColor:{default:u}},setup(r){const t=r;async function d(l,s){const e=l.currentTarget;let o=!1;try{const a=t.getText?await t.getText():t.text;o=await s(a)}catch{o=!1}finally{const a=o?t.tooltipSuccessText:t.tooltipFailText;e instanceof HTMLButtonElement&&(e.setAttribute("data-tooltip-copy-success",String(o)),e.setAttribute("data-tooltip-text",a),window.setTimeout(function(){e instanceof HTMLButtonElement&&e.removeAttribute("data-tooltip-text")},1500))}}return(l,s)=>{const e=n("KButton"),o=n("KClipboardProvider");return f(),y(o,null,{default:i(({copyToClipboard:a})=>[c(e,h(l.$attrs,{appearance:"tertiary",class:["copy-button",{"non-visual-button":!t.hasBorder}],"data-testid":"copy-button",title:t.hideTitle?void 0:t.copyText,type:"button",onClick:p=>d(p,a)}),{default:i(()=>[c(T(x),{color:t.iconColor,title:t.hideTitle?void 0:t.copyText,"hide-title":t.hideTitle},null,8,["color","title","hide-title"]),C(),m(l.$slots,"default",{},()=>[B("span",g,b(t.copyText),1)],!0)]),_:2},1040,["class","title","onClick"])]),_:3})}}});const k=v(S,[["__scopeId","data-v-48672cb5"]]);export{k as C};