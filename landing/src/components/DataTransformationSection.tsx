const emojis = ["😀", "😁", "😂", "🤣", "😃", "😄", "😅", "😆", "😉", "😊", "😋", "😎", "😍", "😘", "🥰", "😗",
    "😙", "😚", "☺️", "🙂", "🤗", "🤩", "🤔", "🤨", "😐", "😑", "😶", "🙄", "😏", "😣", "😥", "😮",
    "🤐", "😯", "😪", "😫", "😴", "😌", "😛", "😜", "😝", "🤤", "😒", "😓", "😔", "😕", "🙃", "🤑",
    "😲", "☹️", "🙁", "😖", "😞", "😟", "😤", "😢", "😭", "😦", "😧", "😨", "😩", "🤯", "😬", "😰",
    "😱", "🥵", "🥶", "😳", "🤪", "😵"];

export default function DataTransformationSection() {
    return (
        <div className="grid grid-cols-1 md:grid-cols-6 border-x-2 border-x-white/90">
            <div className="hidden md:block col-span-1 border-b-2 border-b-white/90 border-r-2 border-r-white/90"></div>

            <div className="col-span-1 md:col-span-2 border-b-2 border-b-white/90 border-r-2 border-r-white/90 bg-[#1e1e1e] p-6">
                <h3 className="font-sekuya text-xl font-bold text-white/90 mb-4">Original Data</h3>
                <pre className="text-xs md:text-sm text-green-400 font-mono overflow-x-auto">
                    {`[
  {
    "id": "1",
    "email": "honour@example.com",
    "name": "Robinson Honour"
  },
  {
    "id": "2",
    "email": "john@example.com",
    "name": "John Doe"
  }
]`}
                </pre>
            </div>

            <div className="col-span-1 md:col-span-2 border-b-2 border-b-white/90 border-r-2 border-r-white/90 bg-[#0a0a0a] p-6">
                <h3 className="font-sekuya text-xl font-bold text-white/90 mb-4">Emoji Encoded</h3>
                <div className="grid grid-cols-10 gap-1">
                    {emojis.map((emoji, i) => (
                        <div key={i} className="text-2xl text-center hover:scale-110 transition-transform">
                            {emoji}
                        </div>
                    ))}
                </div>
            </div>

            <div className="hidden md:block col-span-1 border-b-2 border-b-white/90"></div>
        </div>
    );
}
