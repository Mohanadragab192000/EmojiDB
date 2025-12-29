import type { Metadata } from "next";
import { Geist, Geist_Mono, Inter, Barlow, Bebas_Neue, Saira, Roboto_Slab } from "next/font/google";
import SmoothScroll from "@/components/SmoothScroll";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

const inter = Inter({
  variable: "--font-inter",
  subsets: ["latin"],
  weight: ["400", "500", "600", "700", "800", "900"],
});

const bartle = Barlow({
  variable: "--font-bartle",
  subsets: ["latin"],
  weight: ["900"],
});

const hegarty = Bebas_Neue({
  variable: "--font-hegarty",
  subsets: ["latin"],
  weight: ["400"],
});

const sekuya = Saira({
  variable: "--font-sekuya",
  subsets: ["latin"],
  weight: ["400", "500", "600", "700", "800", "900"],
});

const robotoSlab = Roboto_Slab({
  variable: "--font-roboto-slab",
  subsets: ["latin"],
  weight: ["400", "500", "600", "700", "800", "900"],
});

export const metadata: Metadata = {
  title: "EmojiDB - Encrypted Emoji-Encoded Database",
  description: "A lightweight, secure database that encrypts your data and encodes it into emojis. Fast queries, simple API, and built-in AES-256 encryption for modern applications.",
  keywords: ["emoji database", "encrypted database", "emoji encoding", "secure database", "AES-256 encryption", "Go database", "TypeScript database", "Node.js database"],
  authors: [{ name: "Robinson Honour", url: "https://robinsonhonour.me" }],
  creator: "Robinson Honour",
  publisher: "Robinson Honour",
  metadataBase: new URL("https://emojidb.pxxl.pro"),
  openGraph: {
    type: "website",
    locale: "en_US",
    url: "https://emojidb.pxxl.pro",
    title: "EmojiDB - Encrypted Emoji-Encoded Database",
    description: "A lightweight, secure database that encrypts your data and encodes it into emojis. Fast queries, simple API, and built-in AES-256 encryption.",
    siteName: "EmojiDB",
    images: [
      {
        url: "/og-image.png",
        width: 1200,
        height: 630,
        alt: "EmojiDB - Encrypted Emoji-Encoded Database",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "EmojiDB - Encrypted Emoji-Encoded Database",
    description: "A lightweight, secure database that encrypts your data and encodes it into emojis. Fast queries, simple API, and built-in AES-256 encryption.",
    creator: "@honour_can_code",
    images: ["/og-image.png"],
  },
  robots: {
    index: true,
    follow: true,
    googleBot: {
      index: true,
      follow: true,
      "max-video-preview": -1,
      "max-image-preview": "large",
      "max-snippet": -1,
    },
  },
  icons: {
    icon: "/favicon.ico",
    shortcut: "/favicon.ico",
    apple: "/apple-touch-icon.png",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} ${inter.variable} ${bartle.variable} ${hegarty.variable} ${sekuya.variable} ${robotoSlab.variable} antialiased`}
      >
        <SmoothScroll>{children}</SmoothScroll>
      </body>
    </html>
  );
}
