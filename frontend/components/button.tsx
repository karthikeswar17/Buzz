export default function Button(props: { name: string; onClick: any, width?: string, height?: string }) {
    let name = props.name
    let onClickCallback = props.onClick
    let width = props.width?props.width:"11"
    let height = props.height?props.height:"11"
    return <button className={`rounded-md h-${height} w-${width} bg-blue-500 text-white  hover:bg-blue-400 px-3`} name={name} onClick={onClickCallback}>{name}</button>
}