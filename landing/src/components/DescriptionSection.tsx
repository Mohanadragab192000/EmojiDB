import DiagonalMesh from "./DiagonalMesh";

export default function DescriptionSection() {
    return (
        <div className="grid grid-cols-1 md:grid-cols-6 border-x-2 border-x-white/90">
            <div className="hidden md:block col-span-1 border-b-2 border-b-white/90 border-r-2 border-r-white/90"></div>
            <div className="col-span-1 md:col-span-3 border-b-2 py-6 md:py-5 border-b-white/90 md:border-r-2 md:border-r-white/90 flex flex-col justify-center px-6 md:px-8">
                <h2 className="font-sekuya text-3xl md:text-4xl lg:text-5xl font-bold text-white/90 mb-3 md:mb-4">
                    Database, but encrypted with emojis
                </h2>
                <p className="text-white/70 text-sm md:text-base leading-relaxed">
                    A lightweight, secure database that encrypts your data and encodes it into emojis.
                    Fast queries, simple API, and built-in encryption make it perfect for modern applications.
                </p>
            </div>
            <div className="hidden md:flex col-span-1 flex-col items-center justify-center text-center border-b-2 border-b-white/90 border-r-2 border-r-white/90"></div>
            <div className="hidden md:block col-span-1 border-b-2 border-b-white/90">
                <div className="h-full border-t-[12px] border-r-[12px] border-t-[#4d4d4d] border-r-[#4d4d4d]">
                    <DiagonalMesh />
                </div>
            </div>
        </div>
    );
}
