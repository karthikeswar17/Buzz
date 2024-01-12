export default function ChatBubble(props:{message:string, user:string}){
    let isCurrent = props.user=="0e921028-9a3a-41ba-a90d-f48863a2400a"?true:false
    return (
    <div className={`flex ${isCurrent?"flex-row-reverse":"flex-row"}`}>
            <div className="m-2 min-h-11 min-w-11 flex-shrink-0">
                <img className=" rounded-full h-11 w-11" src="data:image/jpeg;base64,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"/>
            </div>
            <div className="gap-1 m-3">
                <div className={`flex ${isCurrent?"flex-row-reverse":"flex-row"}`}>
                    <span className=" font-bold text-xs">{props.user}</span>
                </div>
                <div className=" rounded-lg bg-gray-100 w-auto p-1 px-2 ">
                    <span>{props.message}</span>
                </div>
            </div>
        </div>
        )
}