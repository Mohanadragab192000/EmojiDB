import DiagonalMesh from "./DiagonalMesh";
import { Github, User2 } from "lucide-react";

export default function FooterSection() {
    return (
        <div className="grid grid-cols-1 md:grid-cols-6 border-x-2 border-x-white/90">
            <div className="col-span-1 md:col-span-2 border-y-2 border-y-white/90 md:border-r-2 md:border-r-white/90 flex flex-col px-6 md:px-8 py-6">
                <h3 className="font-sekuya text-2xl font-bold text-white/90 mb-4">Why did I Build This</h3>
                <p className="text-white/70 text-sm leading-relaxed">
                    Honestly? I was bored and thought it would be cool to encrypt data and encode it as emojis.
                    Turns out it actually works pretty well. The full stuff was built under 48 hours.
                </p>
            </div>

            <div className="hidden md:block col-span-1 border-y-2 border-y-white/90 border-r-2 border-r-white/90">
                <div className="h-full border-t-[12px] border-r-[12px] border-t-[#4d4d4d] border-r-[#4d4d4d]">
                    <DiagonalMesh />
                </div>
            </div>

            <div className="col-span-1 border-y-2 border-y-white/90 md:border-r-2 md:border-r-white/90 flex flex-col px-6 py-6">
                <h3 className="font-sekuya text-lg md:text-xl font-bold text-white/90 mb-4">Tech Stack</h3>
                <p className="text-white/70 text-sm">
                    Go, TypeScript, Node.js
                </p>
            </div>

            <div className="col-span-1 border-y-2 border-y-white/90 md:border-r-2 md:border-r-white/90 flex flex-col px-6 py-6">
                <h3 className="font-sekuya text-lg md:text-xl font-bold text-white/90 mb-4">Links</h3>
                <div className="space-y-3">
                    <a
                        href="https://github.com/ikwerre-dev/emojidb"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="flex items-center gap-2 text-white/80 hover:text-white transition-colors"
                    >
                        <Github size={16} />
                        <span className="text-xs">GitHub</span>
                    </a>
                    <a
                        href="https://robinsonhonour.me"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="flex items-center gap-2 text-white/80 hover:text-white transition-colors"
                    >
                        <User2 size={16} />
                        <span className="text-xs">Portfolio</span>
                    </a>
                </div>
            </div>

            <div className="col-span-1 border-y-2 border-y-white/90 flex flex-col px-6 py-6">
                <p className="text-white/60 text-xs mb-2">© 2025 EmojiDB</p>
                <p className="text-white/50 text-xs">Built by Robinson Honour</p>
            </div>
        </div>
    );
}
