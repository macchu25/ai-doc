import React from 'react';
import { Plus, Folder, FileCode, Database, Activity, ArrowUpRight } from 'lucide-react';
import Navbar from '@/components/Navbar';

export default function Home() {
  return (
    <main className="min-h-screen pt-24 pb-12 px-6 lg:px-12 bg-gradient-premium">
      <Navbar />
      
      <div className="max-w-7xl mx-auto">
        <header className="mb-10">
          <div className="flex flex-col md:flex-row md:items-end justify-between gap-4">
            <div>
              <h1 className="text-4xl font-extrabold tracking-tight text-white mb-2">My Projects</h1>
              <p className="text-muted-foreground max-w-2xl">
                Manage your analyzed codebases and technical documentation. Everything runs locally for total privacy.
              </p>
            </div>
            <button className="flex items-center gap-2 bg-primary hover:bg-primary/90 text-white px-5 py-2.5 rounded-xl font-semibold transition-all shadow-lg shadow-primary/20">
              <Plus className="w-5 h-5" />
              New Project
            </button>
          </div>
        </header>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {/* New Project Placeholder */}
          <div className="group relative flex flex-col items-center justify-center p-12 rounded-3xl border-2 border-dashed border-white/10 hover:border-primary/50 transition-all cursor-pointer bg-white/5 hover:bg-white/10">
            <div className="w-16 h-16 rounded-2xl bg-secondary flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
              <Plus className="w-8 h-8 text-primary" />
            </div>
            <h3 className="text-lg font-bold text-white mb-1">Create New</h3>
            <p className="text-sm text-muted-foreground text-center">Upload ZIP or drag-drop source folder</p>
          </div>

          {/* Sample Project 1 */}
          <ProjectCard 
            title="E-Commerce API" 
            framework="Spring Boot" 
            docsCount={12} 
            status="Ready" 
            lastAnalyzed="2 hours ago" 
          />

          {/* Sample Project 2 */}
          <ProjectCard 
            title="Cardiac Alert System" 
            framework="Go / Next.js" 
            docsCount={8} 
            status="Analyzing" 
            lastAnalyzed="Just now" 
          />

          {/* Sample Project 3 */}
          <ProjectCard 
            title="Portfolio Website" 
            framework="Next.js" 
            docsCount={4} 
            status="Ready" 
            lastAnalyzed="3 days ago" 
          />
        </div>
      </div>
    </main>
  );
}

const ProjectCard = ({ title, framework, docsCount, status, lastAnalyzed }: any) => {
  return (
    <div className="glass-card p-6 rounded-3xl hover:translate-y-[-4px] transition-all cursor-pointer group">
      <div className="flex items-start justify-between mb-6">
        <div className="w-12 h-12 rounded-2xl bg-primary/10 border border-primary/20 flex items-center justify-center">
          <Folder className="w-6 h-6 text-primary" />
        </div>
        <div className="flex items-center gap-1.5 px-2 py-1 rounded-full bg-secondary/50 text-[10px] font-bold uppercase tracking-wider text-muted-foreground border border-white/5">
          <Activity className="w-3 h-3 text-emerald-500" />
          {status}
        </div>
      </div>
      
      <h3 className="text-xl font-bold text-white mb-1 group-hover:text-primary transition-colors flex items-center gap-2">
        {title}
        <ArrowUpRight className="w-4 h-4 opacity-0 group-hover:opacity-100 transition-opacity" />
      </h3>
      <p className="text-sm text-muted-foreground mb-6 flex items-center gap-2">
        <FileCode className="w-4 h-4" />
        {framework}
      </p>

      <div className="pt-6 border-t border-white/5 flex items-center justify-between">
        <div className="flex items-center gap-4">
          <div className="flex flex-col">
            <span className="text-[10px] uppercase font-bold text-muted-foreground tracking-widest">Docs</span>
            <span className="text-white font-bold">{docsCount}</span>
          </div>
          <div className="flex flex-col">
            <span className="text-[10px] uppercase font-bold text-muted-foreground tracking-widest">Last Run</span>
            <span className="text-white font-bold">{lastAnalyzed}</span>
          </div>
        </div>
        <div className="flex -space-x-2">
          {[1, 2, 3].map(i => (
            <div key={i} className="w-7 h-7 rounded-full bg-secondary border-2 border-[#111114] flex items-center justify-center">
              <Database className="w-3 h-3 text-muted-foreground" />
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
