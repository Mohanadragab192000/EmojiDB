export default function StatsSection() {
    return (
        <div className="relative border-x-2 border-x-white/90">
            <div className="absolute inset-0">
                <svg width="100%" height="100%" aria-hidden="true">
                    <defs>
                        <pattern viewBox="0 0 10 10" width="10" height="10" patternUnits="userSpaceOnUse" id="_r12R_6_">
                            <circle cx="5" cy="5" r="1" fill="currentColor" className="fill-white/30"></circle>
                        </pattern>
                    </defs>
                    <rect width="100%" height="100%" fill="url(#_r12R_6_)"></rect>
                </svg>
            </div>

            <div className="relative z-10 p-4 md:p-[3rem]">
                <div className="grid grid-cols-1 md:grid-cols-3 gap-4 md:gap-0 bg-[#0a0a0a]">
                    {[
                        { value: "10,994", label: "Lines of Code" },
                        { value: "40", label: "Code Files" },
                        { value: "48h", label: "Build Time" }
                    ].map((stat, i) => (
                        <div key={i} className="col-span-1 flex flex-col items-center justify-center py-8 md:py-12">
                            <p className="font-sekuya text-3xl md:text-4xl font-bold text-white/90 mb-2">{stat.value}</p>
                            <p className="text-white/60 text-sm">{stat.label}</p>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}
