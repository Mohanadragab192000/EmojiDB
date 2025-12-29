"use client";

import DecorativeBorder from "@/components/DecorativeBorder";
import HeroSection from "@/components/HeroSection";
import DescriptionSection from "@/components/DescriptionSection";
import HowItWorksSection from "@/components/HowItWorksSection";
import DataTransformationSection from "@/components/DataTransformationSection";
import StatsSection from "@/components/StatsSection";
import BenchmarksSection from "@/components/BenchmarksSection";
import FooterSection from "@/components/FooterSection";

export default function Home() {
  const copyToClipboard = () => {
    navigator.clipboard.writeText('npm install @ikwerre-dev/emojidb');
  };

  return (
    <div className="flex min-h-screen lg:px-5 flex-col bg-[#0a0a0a] font-sans">
      <DecorativeBorder side="left" />
      <DecorativeBorder side="right" />

      <div className="px-4 sm:px-8 lg:px-16 min-h-screen flex lg:py-[1rem] flex-col w-full">
        <HeroSection onCopy={copyToClipboard} />
        <DescriptionSection />
        <HowItWorksSection />
        <DataTransformationSection />
        <StatsSection />
        <BenchmarksSection />
        <FooterSection />
      </div>
    </div>
  );
}
