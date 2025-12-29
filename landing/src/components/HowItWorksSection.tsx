import DiagonalMesh from "./DiagonalMesh";
import CodeEditor from "./CodeEditor";

export default function HowItWorksSection() {
    return (
        <div className="grid grid-cols-1 md:grid-cols-6 border-x-2 pt-[.5px] border-x-white/90">
            <div className="hidden md:block col-span-1 h-[30rem] border-b-2 border-b-white/90 border-r-2 border-r-white/90">
                <div className="h-full border-t-[12px] border-r-[12px] border-t-[#4d4d4d] border-r-[#4d4d4d]">
                    <DiagonalMesh />
                </div>
            </div>
            <div className="col-span-1 md:col-span-3 h-[20rem] md:h-[30rem] border-b-2 border-b-white/90 md:border-r-2 md:border-r-white/90 relative">
                <CodeEditor />
            </div>
            <div className="col-span-1 md:col-span-2 min-h-[20rem] md:h-[30rem] border-b-2 border-b-white/90 flex flex-col px-6 md:px-8 py-6">
                <h3 className="font-sekuya text-2xl md:text-3xl font-bold text-white/90 mb-4 md:mb-6">How it Works</h3>
                <div className="space-y-4">
                    {[
                        { title: "Insert Data", desc: "Your data is received via the simple API" },
                        { title: "Encrypt in Memory", desc: "Data is encrypted using AES-256 with your secret key" },
                        { title: "Convert to Bytes", desc: "Encrypted data is converted to byte sequences" },
                        { title: "Encode to Emoji Pairs", desc: "Each byte pair is mapped to a unique emoji combination" },
                        { title: "Store & Query", desc: "Emoji-encoded data is stored for fast retrieval" }
                    ].map((step, i) => (
                        <div key={i} className="flex gap-3">
                            <div className="flex-shrink-0 items-center text-sm font-bold">-</div>
                            <div>
                                <h4 className="font-semibold text-white/90 text-base mb-0.5">{step.title}</h4>
                                <p className="text-sm text-white/60">{step.desc}</p>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}
