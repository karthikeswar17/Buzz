export default function TextBox(props:{name: string, type?: string,onChange: any, width?: string, height?: string}){
    let name= props.name
    let onChange= props.onChange
    let type = props.type
    let width = props.width?props.width:"11"
    let height = props.height?props.height:"11"
    return <input className={`h-${height} w-${width} border-2 border-gray-100 rounded-md box-border w-80 px-3 focus:outline-none focus:bg-gray-100 `} placeholder={name} onChange={onChange} type={type?type:"text"}></input>
}