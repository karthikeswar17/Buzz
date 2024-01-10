

export default function MessageLayout({ children }: { children: React.ReactNode }) {
    return (
        <div className="px-3">
            <div className=" h-14 flex items-center ">
                <span className=" text-lg font-bold">Header</span>
                
            </div>
            <div className="border-b-2"></div>
            <div className="py-3">
                {children}
            </div>
        </div>
    )
}