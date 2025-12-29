const benchmarks = [
    { name: "Open Database", time: "0.720ms" },
    { name: "Define 2 Schemas", time: "9.456ms" },
    { name: "Ingest 1500 Rows", time: "9.417ms" },
    { name: "Unique Constraint Test", time: "0.014ms" },
    { name: "Bulk Update (50)", time: "4.518ms" },
    { name: "Single Update (5)", time: "3.063ms" },
    { name: "Bulk Delete (50)", time: "4.928ms" },
    { name: "Single Delete (5)", time: "3.979ms" },
    { name: "Multi-Table Append", time: "0.018ms" },
    { name: "Flush to Disk", time: "4.043ms" },
    { name: "Inspect File", time: "0.150ms" },
    { name: "Schema Evolution", time: "4.853ms" },
    { name: "Execute Query", time: "4.913ms" },
    { name: "JSON Export", time: "3.153ms" },
    { name: "Generate Secure PEM", time: "0.461ms" },
    { name: "Rotate Master Key", time: "28.313ms" }
];

export default function BenchmarksSection() {
    return (
        <>
            <div className="grid grid-cols-1 border-x-2 border-x-white/90">
                <div className="border-b-2 px-5 border-b-white/90 flex items-center justify-center py-8">
                    <h2 className="font-sekuya text-4xl font-bold text-white/90">Performance Benchmarks</h2>
                </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-4 border-x-2 border-x-white/90">
                {benchmarks.map((benchmark, i) => (
                    <div key={i} className="col-span-1 border-b-2 border-b-white/90 border-r-2 border-r-white/90 px-6 py-4 hover:bg-white/5 transition-colors">
                        <p className="text-white/60 text-xs mb-1">{benchmark.name}</p>
                        <p className="font-sekuya text-xl font-bold text-white/90">{benchmark.time}</p>
                    </div>
                ))}
            </div>

            <div className="grid grid-cols-1 border-x-2 border-x-white/90">
                <div className="border-b-2 border-b-white/90 flex items-center justify-center py-6 bg-white/5">
                    <p className="font-sekuya text-2xl font-bold text-white/90">Total Time: <span className="text-green-400">79.946ms</span></p>
                </div>
            </div>
        </>
    );
}
